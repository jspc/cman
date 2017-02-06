cman
==

`cman` is a PoC config manager thing.

Given a recipe, as per [compiler/fixtures/webserver.yml](compiler/fixtures/webserver.yml), turn into some shell script.

This yml is, actually, any valid [`text/template`](https://golang.org/pkg/text/template/) based file (with the addition of functions from [github.com/moul/funcmap](github.com/moul/funcmap) that compiles into yaml.

This final script will be, should I bother completing the tool, be pushed into some storage to be brought down by an agent on boot of a machine.

It will do idempotence, of course, but the goal has always been to facilitate immutable infrastructure first.


| who       | what |
|-----------|------|
| dockerhub | https://hub.docker.com/r/jspc/cman/   |
| circleci  | https://circleci.com/gh/jspc/cman   |
| licence   | MIT   |


Licence
--

MIT License

Copyright (c) 2017 jspc

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
