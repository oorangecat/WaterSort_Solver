# WaterSort solver

The goal is to build a solver for the "WaterSort" type of game, in which the player has to sort a variable number of vials with 4 or more layer of mixed colors, using 2 additional empty vials (or more, if you got enough real-world money).

WaterSort can me modelled as a blind labyrinth in which each possible color swap is a path turn. Shortest-path algorithms cannot be used as they need a fully generated state-graph. In this case, generating a full state graph would mean solving the whole problem with a breadth-first generation of all the states. 

The initial idea is to develop an heuristic that can help the algorithm predict a favourable state change and explore that path, eventually backtracking if a wrong solution was taken.

The full path is kept to allow backtracking, and a set of discarded nodes is used to avoid re-exploring wrong paths. If the same state is reached by a different path, only the shortest is kept. 
Exploration continues always from the currently-best next node in any parallel path, in an A* fashion.

Any solution that does not imply a full-state exploration will be considered okay-ish.