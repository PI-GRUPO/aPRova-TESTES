FROM alpine:3.19

LABEL org.opencontainers.image.description="Cenário de análise integrada do aPRova"

WORKDIR /app
COPY . .

CMD ["echo", "repositorio-alvo de testes do aPRova"]
