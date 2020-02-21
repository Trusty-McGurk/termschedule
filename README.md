# termschedule
Final project for CPRE431.

Uses a terminal GUI to simulate the RMS and EDF scheduling algorithms on a single processor.
The simulator is controlled using a simple command line, as well as other keyboard inputs. A complete list of inputs is as follows:

`add X Y Z`
- Adds a task to the set with name X, computation time Y, and period Z.

`del X`
- Deletes a task with task name X from the set.

<Enter>
- Run a command.

<TAB>
- Changes the algorithm between RMS and EDF.

<Control-e>
- Changes the right information window between the task info and the RMS exact analysis results.

<UpArrowKey>/<DownArrowKey>
- If there are more tasks than can fit on the screen, scroll through them.

<RightArrowKey>/<LeftArrowKey>
- Scroll forward and backward through time in the schedule.
  
 <Control-c>
- Exit the program.
