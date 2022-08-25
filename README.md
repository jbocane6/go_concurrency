CONCURRENCIAS

Sincronizando acceso a memoria
(bajo nivel o tradicional)
    Sync Package
        Wait Group (Asegurar gorutinas finalicen correctamente)
        Mutex (solo una gorutina puede acceder a una variable a la vez)

Compartiendo la memoria por comunicación
(Alto nivel - Comunicación de procesos secuenciales CSP)
    Channels (Comunicación entre gorutinas)
        Range
        Close
        Select

PROBLEMAS

Data race
2 variables acceden a la misma variable al mismo tiempo y al menos una de ellos es de escritura.

Race condition
2 o más operaciones deben ejecutarse en el orden correcto, pero el programa no se ha escrito para garantizar que se mantenga ese orden.

Deadlock
Un grupo de gorutinas se seperan unas a otras y ninguna de ellas puede continuar

LiveLock
Programa realiza operaciones concurrentes de forma activa, pero no hacen nada para hacer avanzar el estado del programa.

Starvation
Un proceso concurrente no puede obtener los recursos que necesita para realizar su trabajo y no puede progresar.


Exclusión Mutua (Mutex)
Evita que entre más de un proceso a la vez en sección crítica.

Sección Crítica
Fragmento de código donde puede modifir¿carse un recurso compartido.

