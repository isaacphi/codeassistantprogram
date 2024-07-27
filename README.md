# cap code assistant program

This is a simple CLI tool for interacting with LLMs using branching threads. WIP

## Usage

```bash
cap -h
cap thread -h
# etc...

cap thread create myThread
echo "Hi, how are you?" | cap
cap  # Call directly for interactive view
cap thread view
cap thread branch myThread2
cap thread use myThread
cap thread delete myThread
```