package model

import "github.com/luizvnasc/cartola/cartola-commons/util"

type Atleta struct {
	AtletaID         int     `json:"atleta_id"`
	ClubeID          int     `json:"clube_id"`
	PosicaoID        int     `json:"posicao_id"`
	StatusID         int     `json:"status_id"`
	Pontos           float32 `json:"pontos_num"`
	Preco            float32 `json:"preco_num"`
	Media            float32 `json:"media_num"`
	Jogos            int     `json:"jogos_num"`
	Variacao         float32 `json:"variacao_num"`
	Slug             string  `json:"slug"`
	Apelido          string  `json:"apelido"`
	ApelidoAbreviado string  `json:"apelido_abreviado"`
	Nome             string  `json:"nome"`
	Foto             string  `json:"foto"`
	Scout            Scout   `json:"scout"`
}

type Clube struct {
	ID           int    `json:"id"`
	Nome         string `json:"nome"`
	Abreviacao   string `json:"abreviacao"`
	Escudo       Imagem `json:"escudo"`
	NomeFantasia string `json:"nome_fantasia"`
}

// Fechamento representa a data de fechamento do mercado
type Fechamento struct {
	Dia       int    `json:"dia"`
	Mes       int    `json:"mes"`
	Ano       int    `json:"ano"`
	Hora      string `json:"hora"`
	Minuto    string `json:"minuto"`
	Timestamp int    `json:"timestamp"`
}

type Imagem struct {
	X60 string `json:"60x60"`
	X45 string `json:"45x45"`
	X30 string `json:"30x30"`
}

type Mercado struct {
	Rodada    int                      `json:"rodada"`
	Temporada int                      `json:"temporada"`
	Atletas   []Atleta                 `json:"atletas"`
	Clubes    map[string]*Clube        `json:"clubes"`
	Posicoes  map[string]*Posicao      `json:"posicoes"`
	Status    map[string]*StatusAtleta `json:"status"`
}

func NewMercadoFromMap(mp map[string]interface{}) (*Mercado, error) {
	var m = Mercado{}
	err := util.MapToStruct(mp, m)
	return &m, err
}

func (m *Mercado) ToMap() (map[string]interface{}, error) {
	return util.StructToMap(m)
}

type Posicao struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Abreviacao string `json:"abreviacao"`
}

type Scout struct {
	Rodada    int `json:"rodada"`
	Temporada int `json:"temporada"`
	AtletaID  int `json:"atleta_id"`
	//Scouts positivos
	DS float32 `json:"DS"`
	G  float32 `json:"G"`
	A  float32 `json:"A"`
	SG float32 `json:"SG"`
	FS float32 `json:"FS"`
	FF float32 `json:"FF"`
	FD float32 `json:"FD"`
	FT float32 `json:"FT"`
	PS float32 `json:"PS"`
	DE float32 `json:"DE"`
	DP float32 `json:"DP"`
	//Scouts negativos
	GC float32 `json:"GC"`
	CV float32 `json:"CV"`
	CA float32 `json:"CA"`
	GS float32 `json:"GS"`
	PP float32 `json:"PP"`
	PC float32 `json:"PC"`
	FC float32 `json:"FC"`
	I  float32 `json:"I"`
	PI float32 `json:"PI"`
}

// Status do jogo CartolaFC
type Status struct {
	RodadaAtual   int        `json:"rodada_atual"`
	StatusMercado int        `json:"status_mercado"`
	Temporada     int        `json:"temporada"`
	GameOver      bool       `json:"game_over"`
	Fechamento    Fechamento `json:"fechamento"`
}

type StatusAtleta struct {
	Nome string `json:"nome"`
	ID   int    `json:"id"`
}
