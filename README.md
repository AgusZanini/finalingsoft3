# Proyecto final Ingeniería de Software III

## Objetivo: El objetivo proyecto es automatizar la construccion y el despliegue de una aplicacion, asi como las pruebas unitarias y de integracion.

**Desarrollo**: Para el desarrollo del proyecto se uso una aplicacion simple de desarrollo propio que tiene dos funcionalidades:
- Multiplicar dos numeros, mostrar el resultado y almacenar la operacion
- Buscar por id una operacion almacenada y mostrarla.

La aplicacion esta constituida por tres partes fundamentales:
  Backend: Golang, tiene 2 endpoints:
    - `/operations` Con un POST recibe un json donde van a estar los dos numeros a multiplicar, va almacenar la operacion y devolver el resultado para que el frontend lo muestre.
    - `/operations/:id` Con un GET recibe un id y va a traer la operacion que tenga ese id si es que existe.
  Frontend: React.js, se encarga de hacer las peticiones necesarias al backend, de mostrar el resultado de nuevas operaciones y de mostrar las operaciones buscadas por su id.
  Base de datos: Mysql, tiene una tabla `operaciones` que tendra los dos numeros implicados en cada operacion, el resultado y su id correspondiente.

Para el desarrollo del proceso de CI/CD, se opto por usar github actions y se divide en dos worflows:
- Build and publish: Este worflow se corre cada vez que se haga un commit nuevo en la rama main, se encarga de construir la aplicacion, correr los test unitarios y si estos pasan se publican las imagenes en dockerhub
- Deploy: Este workflow se corre manualmente, se encarga de hacer el deploy de los servicios y de la base de datos en google cloud, una vez que se hizo el deploy se corren los test de integracion hechos con codeceptjs
