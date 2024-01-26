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

## Como testar?

Execute o seguinte comando a partir do diretório raíz?

```shell
go-weather-city$ go run cmd/main.go cmd/wire_gen.go
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