package hunt_test

import (
	"github.com/stretchr/testify/require"
	hunt "testdoubles"
	"testing"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		const (
			isSharkHungry = true
			isSharkTired  = false
			sharkSpeed    = 50.0

			tunaName  = "POLA"
			tunaSpeed = 10.2
		)

		shark := hunt.NewWhiteShark(isSharkHungry, isSharkTired, sharkSpeed)

		tuna := hunt.NewTuna(tunaName, tunaSpeed)

		err := shark.Hunt(tuna)

		//assert
		require.NoError(t, err)        // hará que la prueba falle inmediatamente si el err devuelto por el método Hunt no es nil.
		require.False(t, shark.Hungry) // hará que la prueba falle inmediatamente si el campo Hungry del tiburón no es False (es decir, el tiburón todavía tiene hambre después de cazar).
		require.True(t, shark.Tired)   // hará que la prueba falle inmediatamente si el campo Tired del tiburón no es True (es decir, el tiburón no está cansado después de cazar).
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		const (
			isSharkHungry = false
			isSharkTired  = false
			sharkSpeed    = 100.0

			tunaName  = "POLA"
			tunaSpeed = 50.0
		)

		whiteShark := hunt.NewWhiteShark(isSharkHungry, isSharkTired, sharkSpeed)
		tuna := hunt.NewTuna(tunaName, tunaSpeed)

		err := whiteShark.Hunt(tuna)

		// assert
		require.ErrorIs(t, err, hunt.ErrSharkIsNotHungry)            // está comprobando si el error err, devuelto por la función que está siendo probada, es exactamente el mismo error que hunt.ErrSharkIsNotHungry. Esta verificación es precisa y no se basa en el mensaje del error, sino en la igualdad exacta de los errores.
		require.EqualError(t, err, hunt.ErrSharkIsNotHungry.Error()) // valida que el mensaje del error sea el esperado
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		const (
			isSharkHungry = true
			isSharkTired  = true
			sharkSpeed    = 100.0

			tunaName  = "POLA"
			tunaSpeed = 50.0
		)

		whiteShark := hunt.NewWhiteShark(isSharkHungry, isSharkTired, sharkSpeed)
		tuna := hunt.NewTuna(tunaName, tunaSpeed)

		err := whiteShark.Hunt(tuna)

		// assert
		require.ErrorIs(t, err, hunt.ErrSharkIsTired)
		require.EqualError(t, err, hunt.ErrSharkIsTired.Error())
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		const (
			isSharkHungry = true
			isSharkTired  = false
			sharkSpeed    = 50.0

			tunaName  = "POLA"
			tunaSpeed = 100.0
		)

		whiteShark := hunt.NewWhiteShark(isSharkHungry, isSharkTired, sharkSpeed)
		tuna := hunt.NewTuna(tunaName, tunaSpeed)

		err := whiteShark.Hunt(tuna)

		// assert
		require.ErrorIs(t, err, hunt.ErrSharkIsSlower)
		require.EqualError(t, err, hunt.ErrSharkIsSlower.Error())
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		const (
			isSharkHungry = true
			isSharkTired  = false
			sharkSpeed    = 100.0
		)

		whiteShark := hunt.NewWhiteShark(isSharkHungry, isSharkTired, sharkSpeed)
		err := whiteShark.Hunt(nil)

		// assert
		require.ErrorIs(t, err, hunt.ErrTunaIsNil)
		require.EqualError(t, err, hunt.ErrTunaIsNil.Error())
	})

}
