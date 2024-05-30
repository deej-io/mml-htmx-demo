# SSR MML with HTMX Demo

### CAVEATS:
- This is a rough PoC, I make no guarantees about its quality.
- This is HTMX taken to the extreme and is not meant to any best practices - it is
of an exploration of what can be done with SSR mml.

This is an experiment to try run an networked DOM that is controlled entirely
using server-side rendered HTML and HTMX for interactivity.

It currently requires a patched version of JSDOM to fix the
XPathExpression.evaluate implementation
(https://github.com/jsdom/jsdom/pull/3719).

## Requirements

- Node.js
- npm
- go

## Running the servers

Start the API server:

```bash
cd api
go run github.com/cosmtrek/air@latest
```

Start the MML server:
```bash
cd mml-server
npm start
```

Open client endpoint in your browser

http://localhost:808k
