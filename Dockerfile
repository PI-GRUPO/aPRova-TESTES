FROM alpine:3.19

LABEL org.opencontainers.image.title="aPRova - cenário de aprovação manual"

WORKDIR /app
COPY . .

CMD ["echo", "repositorio-alvo de testes do aPRova"]
