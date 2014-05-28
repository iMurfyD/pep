pep
===

A File Compression Program

I'm building this purely for hobby purposes

The idea is that my program losslessly compresses files (I'm going to start with simple text files)

It does it in the standard way of:
  - Taking the file
  - Looking for repitition
  - Represented each repeated block with some shorter symbol which is somehow mapped back to the original memory
  
BUT, what mine (hopefully) does differently is that it concurrently goes throught the file before compression 
and checks to see what the best block size would be for compression, instead of relying on different techniques 
for different data. Maybe, I could expand this into it "learns" how to compress the data before it compresses it.

