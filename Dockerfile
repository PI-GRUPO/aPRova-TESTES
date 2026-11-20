FROM alpine:3.19

LABEL org.opencontainers.image.description="Cenario de teste de risco alto do aPRova"

WORKDIR /app
COPY . .

CMD ["echo", "repositorio-alvo de testes do aPRova"]
