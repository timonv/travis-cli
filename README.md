# Travis CLI

A platform independent Travis status fetcher. No more need to go to the
web or mail to check your build status.

There's a similar ruby gem with similar features. However, this is
platform independent.

## Installing

For now you'll have to clone the source, compile and install using make.

```bash
  $ git clone https://github.com/timonv/travis-cli.git 
  $ make 
  $ make install
```

Or, if you'd prefer a soft link, you can use ```make install soft```.

## Usage

Run ```travis_cli``` in your repo of choice. It will pick the correct
repo, owner and branch automagically. Optionally you can provide these
as params.

Note that if either repo or owner params are nil, it will detect BOTH.

## Future features

tbd.

## Contributing

1. Fork
3. Test
2. Code
3. Test
4. Test the actual cli
5. PR

## License

Copyright (c) 2012 Timon Vonk

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the "Software"), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
