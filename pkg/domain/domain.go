package domain

import "time"

// Priority representa o nível de severidade para os logs
type Priority int

// Níveis de severidade baseados no syslog.h
const (
	LOG_EMERG   Priority = iota // 0
	LOG_ALERT                   // 1
	LOG_CRIT                    // 2
	LOG_ERR                     // 3
	LOG_WARNING                 // 4
	LOG_NOTICE                  // 5
	LOG_INFO                    // 6
	LOG_DEBUG                   // 7
)

// Mapa para converter Priority para string
var PriorityToString = map[Priority]string{
	LOG_EMERG:   "LOG_EMERG",
	LOG_ALERT:   "LOG_ALERT",
	LOG_CRIT:    "LOG_CRIT",
	LOG_ERR:     "LOG_ERR",
	LOG_WARNING: "LOG_WARNING",
	LOG_NOTICE:  "LOG_NOTICE",
	LOG_INFO:    "LOG_INFO",
	LOG_DEBUG:   "LOG_DEBUG",
}

// BaseLog representa os campos comuns a todos os tipos de logs
type BaseLog struct {
	Priority            Priority  // Prioridade do log
	LogLevel            string    // Nível do log (debug, info, error, etc.)
	Timestamp           time.Time // Momento em que o log foi registrado
	GenericErrorMessage string    // Mensagem genérica do log
	ErrorMessage        string    // Mensagem de erro específica
}

// FunctionLog representa um registro de log relacionado a uma função específica
type FunctionLog struct {
	BaseLog             // Incorpora os campos de BaseLog
	FunctionName string // Nome da função registrada
}

// DatabaseLog representa um registro de log relacionado a operações de banco de dados
type DatabaseLog struct {
	BaseLog          // Incorpora os campos de BaseLog
	TableName string // Nome da tabela relacionada à operação
	Query     string // Query executada no banco de dados
}

// RequestLog representa um registro de log relacionado a solicitações (requests)
type RequestLog struct {
	BaseLog             // Incorpora os campos de BaseLog
	Method       string // Método HTTP da solicitação (GET, POST, etc.)
	StatusCode   int    // Código de status HTTP
	Path         string // Caminho da solicitação
	ResponseSize int    // Tamanho da resposta
}
