# Push-Swap Go Implementation

## Project Overview
This project is a Go implementation of the Push-Swap algorithm, a sorting problem that involves sorting a stack of integers using a limited set of operations. The goal is to sort the stack with the minimum number of operations, while adhering to the following constraints:

- You can only use two stacks: `a` (which contains the unsorted numbers) and `b` (which is initially empty)
- The only allowed operations are:
  - `sa`: Swap the first two elements of stack a
  - `sb`: Swap the first two elements of stack b
  - `ss`: Perform sa and sb at the same time
  - `pa`: Push the top element of stack b to stack a
  - `pb`: Push the top element of stack a to stack b
  - `ra`: Rotate stack a (move the top element to the bottom)
  - `rb`: Rotate stack b (move the top element to the bottom)
  - `rr`: Perform ra and rb at the same time
  - `rra`: Reverse rotate stack a (move the bottom element to the top)
  - `rrb`: Reverse rotate stack b (move the bottom element to the top)
  - `rrr`: Perform rra and rrb at the same time

The implementation aims to solve the problem efficiently by minimizing the number of operations required to sort the stack.

## Project Requirements
The goal of the project is to implement an optimized sorting algorithm for a stack using the aforementioned operations. The main focus is on achieving minimal operations while adhering to the constraints.

### Key Features:
- **Efficient Sorting**: The project prioritizes minimizing the number of operations required to sort the stack
- **Cheapest Node Identification**: The algorithm identifies the node with the lowest cost, which represents the least number of operations needed to move it from stack a to stack b
- **Optimal Operations**: The implementation utilizes synchronized rotations (e.g., RotateBoth, RevRotateBoth) to reduce the total number of operations

## Credits
This project would not have been possible without the contributions of the following resources:

- **Yigit Ogun**: Yigit Ogun's excellent article on Medium provided a thorough explanation of the Push-Swap algorithm. His clear, step-by-step guide helped us understand the core concepts of the Push-Swap problem, which were critical for implementing the solution in Go.

- **Thuggonaut (YouTube Channel)**: The Thuggonaut channel on YouTube offered an insightful video explaining the Push-Swap algorithm in C. This video helped enhance our deep understanding of the core concepts and how to translate them into Go, enabling us to successfully implement the project.

## Installation
To use or contribute to this project, follow these steps:

1. Clone the repository to your local machine:
```bash
git clone https://learn.zone01kisumu.ke/git/clomollo/push-swap.git
```

2. Navigate to the project directory:
```bash
cd push-swap
```

3. Build the two main binaries for the `push-swap` and the `checker`:
```bash
go build -o push-swap cmd/push-swap/main.go
go build -o checker cmd/checker/main.go
```

4. To run the project:

* Run the `push-swap` executable with a list of integers to be sorted. Duplicates or non-integers will result in `Error`.

```bash
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
$ ./push-swap "0 one 2 3"
Error
$ ./push-swap
$
```

* Run the `checker` executable with a list of integers and list the operations you expect to sort the integers. The checker will run the operations and display `OK` if the list is sorted, or `KO` if the list is not sorted. You may pipe the output of the `push-swap` directly into `checker` to verify the correctness of the program.

```bash
$ ./checker "3 2 1 0"
sa
rra
pb
KO
$ echo -e "rra\npb\nsa\n" | ./checker "3 2 one 0"
Error
$ echo -e "rra\npb\nsa\nrra\npa"
rra
pb
sa
rra
pa
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
$ ./checker
$
```

## Testing
The project includes unit tests to ensure correct functionality of the implemented algorithms. To run the tests, use the following command:

```bash
go test ./... -v
```

## Usage
The core functionality of the project revolves around the following steps:

1. **Input**: The program accepts a list of integers (stack a) that need to be sorted
2. **Operations**: The program then performs a series of operations (as defined above) to sort the stack
3. **Output**: The output is the list of operations performed to achieve the sorted stack

## Contributing
Feel free to fork this repository, submit issues, or open pull requests. If you have any improvements or suggestions, don't hesitate to contribute!

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
