
# 🗓️CLI TODO APP
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)

This CLI-based ToDo app, developed in Go, enables efficient task management directly from the command line. It provides custom flags to handle essential task operations: adding, editing, and deleting tasks, as well as toggling their completion status. The app is simple yet powerful, helping users maintain and organize their tasks with ease and flexibility.


## Screenshots

![App Screenshot](https://github.com/codesbysagar/TODO/blob/main/screentshot/sampleImg.png)


## Tech Stack

**Language:** GO

**Packages:** aquasecurity/table


## GUIDE to use

Run the below command in directory to clone the TODO app

```bash
  git clone https://github.com/codesbysagar/TODO.git
```
***feel free to explore and make any kind of changes and improvement to the APP***
## Commands

#### Get the current TODO list

```bash
  go run . -list
```

#### Add new Task in TODO list

```bash
  go run . -add <"task name">
```

#### To delete a task from TODO list

```bash
  go run . -del <index> //example -del 0
```

#### To toggle the completion status

```bash
  go run . -toggle <index> //example -toggle 1
```

#### To edit the existing task in TODO list

```bash
  go run . -edit <index:"new task">
```




## License

MIT License

Copyright (c) 2024 **codesbysagar**

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
