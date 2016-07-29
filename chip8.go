package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

const MEMSIZE = 4096

// Structure that represents the cpu
type CPU struct {
  mem [MEMSIZE] uint8 // memory of 8 bits cell
  v[16] uint8 //16 general purpose registers of 8 bits
  
  stack[16] uint16 
  sp uint16 // stack pointer
  
  i uint16 // 16 bits for storage of memory adresses
  dt, st uint8 // delay and sound timers
  pc uint16 // progam counter
}

// Initialize or reset the state of the cpu
func initialize(cpu *CPU) {
  cpu.sp = 0x00
  cpu.i  = 0x00
  cpu.dt = 0x00
  cpu.st = 0x00
  cpu.pc = 0x200 // Set the program counter to 512
  for i := range cpu.mem { cpu.mem[i] = 0 } //Set all memory positions to 0
  for i := 0; i < 16; i++ {
    cpu.stack[i] = 0
    cpu.v[i] = 0
  }
}

func load_rom(cpu *CPU) {
  rom, err := ioutil.ReadFile("./PONG")
  
  if err != nil{
    fmt.Println("Cannot open ROM file")
    os.Exit(1)
  }

  copy(cpu.mem[0x200:], rom) // Load the rom in memory with 512(0x200) offset
  // fmt.Println(rom)

}

func main() {
  var cpu CPU // Create the cpu
  initialize(&cpu)
  load_rom(&cpu) // Load the rom into memory

  mustQuit := false
  var opcode uint16

  var nnn uint16
  var kk uint8
  var n uint8
  var x uint8
  var y uint8
  var p uint8

  for(!mustQuit){
    
    position := cpu.pc // Position of the first byte for the opcode
    // fmt.Println(cpu.pc)
    next_position := cpu.pc + 1 // Position of the second byte for the opcode

    first_byte := uint16(cpu.mem[position]) 
    second_byte := uint16(cpu.mem[next_position])

    opcode = (first_byte << 8) | second_byte // Concat the bytes to for the opcode

    // opcode = cpu.mem[position] | cpu.mem[next_position] 
    if next_position == MEMSIZE - 1 {
      cpu.pc = 0
    } else {
      cpu.pc = next_position + 1// Update the program counter of the cpu
    }

    nnn = opcode & 0x0FFF
    kk = uint8( opcode & 0x0FF )
    n = uint8( opcode & 0xF )
    x = uint8( (opcode >> 8) & 0xF )
    y = uint8( (opcode >> 4) & 0xF )
    p = uint8( (opcode >> 12) )

    switch p {
    case 0:
      if opcode == 0x00E0 {
        fmt.Println("CLS") 
      } else if opcode == 0x00EE {
        fmt.Println("RET") 
      }
    case 1:
      fmt.Println("JP", nnn) // nnn
    case 2:
      fmt.Println("CALL") // nnn
    case 3:
      fmt.Println("SE") // x kk
    case 4:
      fmt.Println("SNE") // x kk
    case 5:
      fmt.Println("SE", x, y) // x y 0
    case 6:
      fmt.Println("LD") // x kk
    case 7:
      fmt.Println("ADD") // x kk
    case 8:
      switch n {
      case 0:
        fmt.Println("LD") // x y 0
      case 1:
        fmt.Println("OR") // x y 1
      case 2:
        fmt.Println("AND") // x y 2
      case 3:
        fmt.Println("XOR") // x y 3
      case 4:
        fmt.Println("ADD") // x y 4
      case 5:
        fmt.Println("SUB") // x y 5
      case 6:
        fmt.Println("SHR") // x y 6
      case 7:
        fmt.Println("SUBN") // x y 7
      case 0xE:
        fmt.Println("SHL") // x y E
      }
    case 0xA:
      fmt.Println("LD") // nnn
    case 0xB:
      fmt.Println("JP") // nnn
    case 0xC:
      fmt.Println("RND") // x kk
    case 0xD:
      fmt.Println("DRW") // x y n
    case 0xE:
      if kk == 0x9E {
        fmt.Println("SKP") // x
      } else if kk == 0xA1 {
        fmt.Println("SKNP") // x
      }
      fmt.Println("SNE") // x y 0
    case 0xF:
      switch kk {
      case 0x07:
        fmt.Println("LD") // x
      case 0x0A:
        fmt.Println("LD") // x
      case 0x15:
        fmt.Println("LD") // x
      case 0x18:
        fmt.Println("LD") // x
      case 0x1E:
        fmt.Println("ADD") // x
      case 0x29:
        fmt.Println("LD") // x
      case 0x33:
        fmt.Println("LD") // x
      case 0x55:
        fmt.Println("LD") // x
      case 0x65:
        fmt.Println("LD") // x
      }
    }

    // fmt.Println(opcode)
  }
}