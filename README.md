# *Bootcamp Go*
# **Desafío Go Web💥**


Objetivo
El objetivo de esta guía práctica es poder afianzar y profundizar los conceptos vistos en Go Web. Para esto, vamos a plantear un desafío integrador que nos permitirá repasar los temas que estudiamos. 

Planteo
Una aerolínea utiliza un programa para calcular diversas informaciones sobre vuelos que ocurren en la misma. Si bien su programa se encuentra funcional y sin dificultades, han decidido convertir el mismo en una API Rest.

¿Are you ready? 


  
Desafío
Como primer objetivo se nos solicita realizar un programa con arquitectura REST, para comenzar tienen la base de proyecto a continuación:

¡Atención! El repository.go ya se encuentra desarrollado.


Comenzamos…

Antes de comenzar a trabajar, crear un repositorio en su cuenta de github con el nombre: “desafio-goweb-nombreapellido”
Una vez finalizado el paso anterior estamos listos para comenzar a trabajar. Recuerden crear el módulo con el nombre de su repositorio de github.com creado previamente.


  go mod init “desafio-goweb-nombreapellido”



Requerimiento 1: 
Crear los packages correspondientes para que la arquitectura de la aplicación cumpla con la definición de REST y el Diseño Orientado a Dominio.
Tip: Los paquetes internal y cmd son fundamentales.

Requerimiento 2: 

Ubicar los archivos de manera tal que cada uno quede en su respectivo package.


Tip: Recordar chequear los imports.

Requerimiento 3: 
Desarrollar los métodos correspondientes a la estructura Ticket. Uno de ellos debe devolver la cantidad de tickets de un destino. El  otro método debe devolver el promedio de personas que viajan a un país determinado en un dia:






Requerimiento 4: 
Una vez desarrollado el servicio y el repositorio, deben desarrollar los controladores para los siguientes endpoints:

GET - “/ticket/getByCountry/:dest”
GET - “/ticket/getAverage/:dest”



Requerimiento 5: 
Correr el main.go y testear los endpoints con la herramienta de Postman.

Tip: El comando go mod tidy puede ser necesario previo a correr el main.



Extra: 
Si terminaste todo lo solicitado previamente y quieres seguir agregando funcionalidad al sistema de tickets de la aerolínea, puedes hacerlo. 
			
