/**
 * ViaCEP JSON GO client test
 */
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Consultar( t *testing.T){
	t.Run("CEP_NAO_ENCONTRADO_404", func(t *testing.T) {
		cep := "7800001"
		_, err := Consultar(cep)
		assert.Error(t, err)
	})
	t.Run("CEP_NAO_ENCONTRADO_ERRO", func(t *testing.T) {
		cep := "000000000"
		_, err := Consultar(cep)
		assert.Error(t, err)
	})
	t.Run("CEP_EXISTENTE", func(t *testing.T) {
		cep := "01001000"
		resultado, _ := Consultar(cep)
		assert.Equal(t, "3550308", resultado.Ibge)
	})
}
