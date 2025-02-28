def romano_a_arabigo(entrada):
    # Definimos los valores de cada letra romana
    valores = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000
    }
    
    total = 0  # Inicializamos el total en 0
    i = 0  # Usamos un índice para recorrer la cadena

    while i < len(entrada):
        # Obtenemos el valor del carácter actual
        valor_actual = valores[entrada[i]]
        
        # Verificamos si hay un siguiente carácter
        if i + 1 < len(entrada):
            valor_siguiente = valores[entrada[i + 1]]
            # Si el valor actual es menor que el siguiente, restamos
            if valor_actual < valor_siguiente:
                total -= valor_actual
            else:
                total += valor_actual
        else:
            # Si no hay siguiente carácter, simplemente sumamos
            total += valor_actual
        
        i += 1  # Pasamos al siguiente carácter

    return total  # Devolvemos el total calculado

def arabigo_a_romano(numero):
    # Definimos los valores de los números romanos en orden descendente
    valores = [
        (1000, 'M'),
        (900, 'CM'),
        (500, 'D'),
        (400, 'CD'),
        (100, 'C'),
        (90, 'XC'),
        (50, 'L'),
        (40, 'XL'),
        (10, 'X'),
        (9, 'IX'),
        (5, 'V'),
        (4, 'IV'),
        (1, 'I')
    ]
    
    resultado = ""  # Inicializamos el resultado como una cadena vacía

    for valor, simbolo in valores:
        # Mientras el número sea mayor o igual al valor, añadimos el símbolo
        while numero >= valor:
            resultado += simbolo
            numero -= valor  # Restamos el valor del número

    return resultado  # Devolvemos el resultado en formato romano

# Ejemplo de uso
entrada_romano = "MMMCMCIX"  # Este es el número romano para 3999
resultado_arabigo = romano_a_arabigo(entrada_romano)
print("El valor de la entrada '{}' es: {}".format(entrada_romano, resultado_arabigo))

entrada_arabigo = 3985  # Este es el número arábigo que queremos convertir
resultado_romano = arabigo_a_romano(entrada_arabigo)
print("El valor de la entrada '{}' en romano es: {}".format(entrada_arabigo, resultado_romano)) # type: ignore