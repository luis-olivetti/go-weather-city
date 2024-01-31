# Go Weather City

## Objetivo
Desenvolver um sistema em Go que receba um CEP, identifique a cidade e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). Este sistema será implantado no Google Cloud Run.

## Requisitos

### Funcionalidades Principais
- Receber um CEP válido de 8 dígitos.
- Pesquisar o CEP para encontrar o nome da localização.
- Retornar as temperaturas formatadas em Celsius, Fahrenheit e Kelvin.

### Respostas Adequadas
- **Em Caso de Sucesso:**
  - Código HTTP: 200
  - Response Body: `{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`

- **Em Caso de Falha, CEP Inválido:**
  - Código HTTP: 422
  - Mensagem: "invalid zipcode"

- **Em Caso de Falha, CEP Não Encontrado:**
  - Código HTTP: 404
  - Mensagem: "cannot found zipcode"

### Dicas
- Utilize a API ViaCEP (ou similar) para encontrar a localização desejada: [ViaCEP](https://viacep.com.br/)
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: [WeatherAPI](https://www.weatherapi.com/)
- Para a conversão de Celsius para Fahrenheit, utilize a fórmula: `F = C * 1,8 + 32`
- Para a conversão de Celsius para Kelvin, utilize a fórmula: `K = C + 273`

## Entrega

### Código-Fonte
- Código-fonte completo da implementação.

### Documentação
- Documentação explicando como rodar o projeto em ambientes de desenvolvimento e produção.

### Testes Automatizados
- Testes automatizados demonstrando o funcionamento.

### Docker/Docker-Compose
- Utilize Docker e Docker-Compose para facilitar os testes da aplicação.

### Google Cloud Run
- Realize o deploy no Google Cloud Run (free tier).
- Forneça o endereço ativo para acesso.

## Como executar?

### Ambiente Dev
Execute o seguinte comando através do Docker Compose:

```shell
go-weather-city$ DOCKERFILE=Dockerfile.dev docker-compose up --build
```

### Ambiente Produção
Execute o seguinte comando através do Docker Compose:

```shell
go-weather-city$ DOCKERFILE=Dockerfile.prod docker-compose up --build
```

### Google Cloud

O projeto também foi publicado no Google Cloud neste endereço:

- https://go-weather-city-zxwp3x2zkq-uc.a.run.app/city?zipCode=47804112

### Observações

No diretório **api** encontra-se os arquivos **http** para facilitar os testes com a extensão **REST Client**  (VSCode).

## Docker

Gerar imagem:

```shell
go-weather-city$ docker build -t goweather .
```

Gerar contâiner mapeado na porta 8080:

```shell
go-weather-city$ docker run -p 8080:8080 goweather
```

## Verificando a cobertura de testes

Simples:
```shell
go-weather-city$ go test -cover ./...
```

Detalhado:
```shell
go-weather-city$ go test -coverprofile=coverage.out ./...
go-weather-city$ go tool cover -html=coverage.out -o coverage.html
```




