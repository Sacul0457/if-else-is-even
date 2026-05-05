FROM oven/bun:canary-alpine AS builder
WORKDIR /app

RUN printf 'console.log(+process.argv[2]%%2===0?"even":"odd")\n' > index.ts \
 && bun build index.ts --compile --minify --outfile is-even

FROM oven/bun:canary-alpine
WORKDIR /app
COPY --from=builder /app/is-even .

CMD ["./is-even"]