# 📚 Estudos de Go com Testes

Este repositório contém uma coleção de pequenos pacotes organizados para estudar conceitos da linguagem Go, seguindo o material disponível no excelente guia [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/criando-uma-aplicacao/introducao).

Cada pacote aborda um conceito ou funcionalidade específica e pode ser executado individualmente com os testes correspondentes. Os testes foram organizados por mim para seguir a estrutura AAA (Arrange, Act, Assert)

## 🧪 Como rodar os testes

Você pode executar os testes de cada pacote usando o comando:

```bash
go test
```

Navegue até o diretório do pacote que deseja testar e execute o comando acima.

**Exemplo:**

```bash
cd internal/hello
go test
```

## 📦 Estrutura do projeto

A estrutura do repositório é a seguinte. Cada diretório representa uma etapa ou um conceito abordado:

```
internal
├── hello/
├── integer/
├── iterator/
├── arrays/
├── pointers/
...
```

## 📖 Material de referência

Todo o estudo foi baseado no guia gratuito e em português:

🔗 [Aprenda Go com Testes - Larien](https://larien.gitbook.io/aprenda-go-com-testes/criando-uma-aplicacao/introducao)

## 💡 Objetivo

O objetivo principal deste repositório é reforçar os conceitos de Go por meio da prática com testes, TDD e exemplos simples e incrementais.
