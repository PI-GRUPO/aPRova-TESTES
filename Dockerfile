FROM alpine:3.19

LABEL br.aprova.cenario="teste-dashboard-20261122"

WORKDIR /app
COPY . .

CMD ["echo", "repositorio-alvo de testes do aPRova"]
