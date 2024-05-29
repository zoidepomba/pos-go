package main

func main() {
	//Memoria -> Endereço -> valor
	a := 10                   // 1. A variável 'a' é declarada e inicializada com o valor 10.
	var ponteiro *int = &a    // 2. Um ponteiro 'ponteiro' para um inteiro é declarado e inicializado com o endereço de 'a'.ponteiro é uma variável do tipo ponteiro para um inteiro (*int). Ela é inicializada com o endereço de a (&a). Isso significa que ponteiro agora aponta para a.
	*ponteiro = 20            // 3. O valor apontado por 'ponteiro' (que é 'a') é alterado para 20. O operador * é usado para acessar o valor na memória apontada por ponteiro. Como ponteiro aponta para a, *ponteiro refere-se ao valor de a. Então, *ponteiro = 20 muda o valor de a para 20.
	b := &a                   // 4. Uma nova variável 'b' é declarada e inicializada com o endereço de 'a'.
	*b = 30                   // 5. O valor apontado por 'b' (que é 'a') é alterado para 30.
	println(a)                // 6. O valor de 'a' é impresso.
}
