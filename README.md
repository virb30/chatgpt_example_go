# Microservice Docker + ChatGPT + Twilio

Microsserviço para integrar ChatGPT ao Whatsapp

## Conceitos OpenAI / ChatGPT

### Tokens

Tokens para ChatGPT são unidades básicas de processamento de texto usadas pelo modelo durante treinamento e geração das respostas

### Models

Versões de modelos de linguagemm treinados usando a arquitetura GPT (Generative Pre-trained Transformer). São redes neurais profundas treinadas em grandes quantidades de texto para aprender a compreender e gerar texto humano coerente

### Funcionamento básico da API do ChatGPT

- Mensagem inicial do sistema
- Pergunta do usuário (user)
- Resposta do chatgpt (assistant)
- Pergunta do usuário (user)
- Resposta do chatgpt (assistant)

- Mensagens vão se acumulando para armazenar contexto (REST - Não guarda estado)
- Quando não couber mais tokens, precisamos remover mensagens para a nova poder entrar

### Tokens e contexto

- O segredo de tudo é fazer a contagem dos tokens
- Sabendo quantos tokens estamos utilizando e a quantidade máxima do modelo, podemos acumular mensagens.
- Quanto mais mensagens, melhor a resposta por causa do contexto


## Microservice ChatGPT

### Arquitetura

- Clean Architecture

### Requisitos

- Coração da aplicação deve ter suas regras de negócio consolidadas
- Coração da aplicação não sabe que existe API da OpenAI
- Armazenar conversações em um banco de dados
- Usuário poderá informar seu "user_id" como referência para ter acesso as conversas de um determinado usuário
- Servidor  Web e gRPC para realizar as conversas
- Precisaremos gerar um Token no site da OpenAI para termos acesso a API
- A autenticação do microsserviço será realizada via token fixo em um arquivo de configuração.


### Tecnologias

- Docker
- Linguagem Go
- MySQL