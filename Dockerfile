FROM alpine:3.19

LABEL br.aprova.cenario="teste-aprovacao-20261124"

WORKDIR /app
COPY . .

CMD ["echo", "repositorio-alvo de testes do aPRova"]
