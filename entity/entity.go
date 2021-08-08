
package entity

//Estructura consumida desde kafka

// EventoPortabilidadSolicitud - CSGS...
type EventoPortabilidadSolicitud struct {
	Key               string         `json:"key"`
	IDSolicitudGestor string         `json:"idSolicitudGestor"`
	Numero            int            `json:"numero"`
	TipoSolicitud     int            `json:"tipoSolicitud"`
	FechaCreacion     string         `json:"fechaCreacion"`
	FechaFin          string         `json:"fechaFin"`
	TipoVerificacion  int            `json:"tipoVerificacion"`
	Cliente           Cliente        `json:"cliente"`
	Portabilidad      []Portabilidad `json:"portabilidad"`
	Preferencias      *Preferencias  `json:"preferencias"`
}

// KeySelector Usado para serializar la key e ignorar el resto
type KeySelector struct {
	Key           string      `json:"key"`
	TipoSolicitud interface{} `json:"tipoSolicitud"`
}

// EventoPortabilidadCertificado - SDSD...
type EventoPortabilidadCertificado struct {
	Key               string `json:"key"`
	IDSolicitudGestor string `json:"idSolicitudGestor"`
	Numero            int    `json:"numero"`
	RutEjecutivo      string
	TipoSolicitud     string       `json:"tipoSolicitud"`
	FechaCreacion     string       `json:"fechaCreacion"`
	FechaFin          string       `json:"fechaFin"`
	TipoVerificacion  int          `json:"tipoVerificacion"`
	Cliente           Cliente      `json:"cliente"`
	Portabilidad      Portabilidad `json:"portabilidad"`
}

// Direcciones  ...
type Direcciones struct {
	Tipo         int    `json:"tipo"`
	Ciudad       string `json:"ciudad"`
	Comuna       string `json:"comuna"`
	Direccion    string `json:"direccion"`
	Departamento string `json:"departamento"`
	Villa        string `json:"villa,omitempty"`
}

// Telefonos ...
type Telefonos struct {
	Tipo   int `json:"tipo"`
	Numero int `json:"numero"`
}

// Rentas ...
type Rentas struct {
	TipoRenta int `json:"tipoRenta"`
	Monto     int `json:"monto"`
}

// Empleo ...
type Empleo struct {
	TipoEmpleo        int    `json:"tipoEmpleo"`
	Actividad         int    `json:"actividad"`
	TipoRenta         int    `json:"tipoRenta"`
	AntiguedadLaboral string `json:"antiguedadLaboral"`
}

// Cliente ...
type Cliente struct {
	Rut                  string        `json:"rut"`
	Nombres              string        `json:"nombres,omitempty"`
	ApellidoPaterno      string        `json:"apellidoPaterno,omitempty"`
	ApellidoMaterno      string        `json:"apellidoMaterno,omitempty"`
	FechaNacimiento      string        `json:"fechaNacimiento,omitempty"`
	Sexo                 string        `json:"sexo,omitempty,omitempty"`
	Nacionalidad         string        `json:"nacionalidad,omitempty"`
	CodigoPaisResidencia string        `json:"codigoPaisResidencia,omitempty"`
	CodigoPaisNacimiento string        `json:"codigoPaisNacimiento,omitempty"`
	CodigoPais           string        `json:"codigoPais,omitempty"`
	Email                string        `json:"email"`
	Direcciones          []Direcciones `json:"direcciones,omitempty"`
	Telefonos            []Telefonos   `json:"telefonos,omitempty"`
	Rentas               []Rentas      `json:"rentas,omitempty"`
	Empleo               Empleo        `json:"empleo,omitempty"`
}

// Documentos ...
type Documentos struct {
	CertificadoDeuda string `json:"certificado_deuda"`
}

// CTAALL ...
type CTAALL struct {
	PortarTodasCuentasBancarias string  `json:"portar_todas_cuentas_bancarias,omitempty"`
	CTACTE                      *CTACTE `json:"CTACTE,omitempty"`
	CTAVIS                      *CTAVIS `json:"CTAVIS,omitempty"`
	CTAAHO                      *CTAAHO `json:"CTAAHO,omitempty"`
}

// ROTALL ...
type ROTALL struct {
	PortarTodosCreditosRotativos string  `json:"portar_todos_creditos_rotativos,omitempty"`
	LINCRE                       *LINCRE `json:"LINCRE,omitempty"`
	TARCRE                       *TARCRE `json:"TARCRE,omitempty"`
}

// CREALL ...
type CREALL struct {
	PortarTodosCreditos string  `json:"portar_todos_creditos,omitempty"`
	CRECON              *CRECON `json:"CRECON,omitempty"`
	CREHIP              *CREHIP `json:"CREHIP,omitempty"`
	CREAUT              *CREAUT `json:"CREAUT,omitempty"`
}

// OtrosProductos ...
type OtrosProductos struct {
	PortarOtrosProductos string `json:"portar_otros_productos"`
}

// Portabilidad ...
type Portabilidad struct {
	Rut                  string         `json:"rut,omitempty"`
	Nombre               string         `json:"nombre,omitempty"`
	Documentos           *Documentos    `json:"documentos,omitempty"`
	PortarTodosProductos string         `json:"portar_todos_productos,omitempty"`
	CTAALL               CTAALL         `json:"CTAALL"`
	ROTALL               ROTALL         `json:"ROTALL"`
	CREALL               CREALL         `json:"CREALL"`
	OtrosProductos       OtrosProductos `json:"otros_productos"`
}

// Preferencias Lista de preferencias
type Preferencias struct {
	CREAUT CREAUT `json:"CREAUT"`
	CRECON CRECON `json:"CRECON"`
	CREHIP CREHIP `json:"CREHIP"`
	CTACTE CTACTE `json:"CTACTE"`
	CTAVIS CTAVIS `json:"CTAVIS"`
	TARCMR TARCMR `json:"TARCMR"`
}

// CRECON Credito de Consumo
type CRECON struct {
	Aplica bool `json:"aplica"`
}

// CREHIP Credito Hipotecario
type CREHIP struct {
	Aplica bool `json:"aplica"`
}

// CREAUT Credito Automotriz
type CREAUT struct {
	Aplica bool `json:"aplica"`
}

// CTACTE Cuenta Corriente
type CTACTE struct {
	Aplica bool `json:"aplica"`
}

// CTAVIS Cuenta Vista
type CTAVIS struct {
	Aplica bool `json:"aplica"`
}

// CTAAHO Cuenta de Ahorros
type CTAAHO struct {
	Aplica bool `json:"aplica"`
}

// LINCRE Linea de Credito
type LINCRE struct {
	Aplica bool `json:"aplica"`
}

// TARCRE Tarjeta de Credito
type TARCRE struct {
	Aplica bool `json:"aplica"`
}

// TARCMR Tarjeta CMR
type TARCMR struct {
	Aplica bool `json:"aplica"`
}

