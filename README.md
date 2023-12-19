# Ejercicio para Validación de [Go - Introducción]

Desarrollar un programa que tenga un listado de Figuras en memoria, permitiendo:\
Listar todas las Figuras\
Obtener una Figura a través de su ID\
Agregar nuevas Figuras

Se deben poder crear Figuras de tipo Elipse, Rectángulo y Triángulo.\
Una Figura se compondrá de un ID que debe ser único y los datos específicos de cada una (base y altura, o alto y largo, o radio a y radio b).\
Cada una de estas Figuras debe tener su propio método para calcular el Área y para hacer un Print de sus datos básicos.

Se puede optar por un main que tenga un menu y lea las instrucciones de teclado, o un main que cargue las figuras desde un archivo y las liste

A tener en cuenta:\
El listado debe ser un map[string]Figura\
Los modelos deben estar en un package separado\
