import unittest
from modulo_conversion import convertir_arabico_a_romano

class TestConversionArabigoRomano(unittest.TestCase):

    def test_conversion_uno(self):
        resultado = convertir_arabico_a_romano(1)
        self.assertEqual(resultado, "I")

    def test_conversion_cinco(self):
        resultado = convertir_arabico_a_romano(5)
        self.assertEqual(resultado, "V")

    def test_conversion_diez(self):
        resultado = convertir_arabico_a_romano(10)
        self.assertEqual(resultado, "X")

    def test_conversion_3999(self):
        resultado = convertir_arabico_a_romano(3999)
        self.assertEqual(resultado, "MMMCMXCIX")

    def test_conversion_fuera_de_rango(self):
        with self.assertRaises(ValueError):
            convertir_arabico_a_romano(4000)

if __name__ == '__main__':
    unittest.main()
