package stm32

import "github.com/qmk/qmk_firmware/runtime/volatile"

type GPIO_Type struct {
	MODER  volatile.Register32
	OTYPER volatile.Register32
	PUPDR  volatile.Register32
	IDR    volatile.Register32
	ODR    volatile.Register32
	BSRR   volatile.Register32
	AFRL   volatile.Register32
	AFRH   volatile.Register32
}

type RCC_Type struct {
	AHB2ENR  volatile.Register32
	APB1ENR1 volatile.Register32
	APB2ENR  volatile.Register32
}

var (
	RCC   = &RCC_Type{}
	GPIOA = &GPIO_Type{}
	GPIOB = &GPIO_Type{}
	SPI1  = &SPI_Type{}
	I2C1  = &I2C_Type{}
	USBFS = &USB_Type{}
)

const (
	RCC_AHB2ENR_GPIOAEN = 1 << 0
	RCC_AHB2ENR_GPIOBEN = 1 << 1

	RCC_APB2ENR_SPI1EN   = 1 << 12
	RCC_APB1ENR1_I2C1EN  = 1 << 21
	RCC_APB1ENR1_USBFSEN = 1 << 24
)

type SPI_Type struct {
	CR1 volatile.Register32
	CR2 volatile.Register32
	SR  volatile.Register32
	DR  volatile.Register32
}

type I2C_Type struct {
	CR1     volatile.Register32
	TIMINGR volatile.Register32
	ISR     volatile.Register32
	TXDR    volatile.Register32
	RXDR    volatile.Register32
}

type USB_Type struct {
	EP0R volatile.Register32
}
