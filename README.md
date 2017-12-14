# DESBreaker
[![Build Status](https://travis-ci.org/mtenrero/DESBreaker.svg?branch=master)](https://travis-ci.org/mtenrero/DESBreaker)
Native program written in go designed to reveal the encryption key (brute-forcing) of a DES encrypted HEX block given the unencrypted message


## Usage
The program is precompiled for all platforms and OSes. Simply download and run your binary with the following parameters:

`desbreaker -hexKnown="HEXEncryptedValue" -plain="Decrypted value of hexKnown" -force="HEX to Decode"`

You also can print a help message if needed: 

`desbreaker -h`