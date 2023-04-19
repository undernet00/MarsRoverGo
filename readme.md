# Mars Rover

You need to code a program to validate the commands that will be sent to a new Rover in Mars.
Each Rover is included in a square and can receive the next commands: Advance (A), Turn left (L), Turn Right (R). 

The program must validate that the Rover will be within the bounds of the square and must indicate the final orientation. 

The program will receive the dimensions of the square (width x height) and can assume that the coordinate (0,0) is the bottom left corner. 
Additionally, will receive initial coordinates of the Rover and its initial orientation (N, S, E, W). Also, it will receive a set of commands like the next one; “AALAARALA”. There is not fixed limit of number of input commands. It can be assumed that there are not obstacles in the square. 

The program must validate that all the commands can be executed without being out of the defined bounds and must also return True or False indicating if the commands are valid. Additionally, it must return the orientation and final coordinates of the Rover As example: True, N, (4,5). 

The source code delivered for this test must include everything needed to run the solution and get the proper results. Must be compiled without any errors. Any language is permitted. 

**Tips** 

* This test will be used to check technical skills so consider all aspects you use to use every day to keep quality and architecture and everything else that define you as the good professional you are. 

* Be honest with the time you used to complete the whole test. Good luck!

**Dev Assumptions**

* In a scenario where the list of commands will leave the rover out of the map. The rover will move to the last valid position. And the program will return false to state that the list of commands are not valid. 