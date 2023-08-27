# WaterSort solver

The goal is to build a solver for the "WaterSort" type of game, in which the player has to sort a variable number of vials with 4 or more layer of mixed colors, using 2 additional empty vials (or more, if you got enough real-world money).

The initial idea is to develop an heuristic that can help the algorithm predict a favourable state change and explore that path, eventually backtracking if a wrong solution was taken. 

Any solution that does not imply a full-state exploration will be considered okay-ish.

