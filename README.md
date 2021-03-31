# Descripción

Esta es una api dummy para usar de ejemplo en la clase de desarrollo de aplicaciones en la nube. La api solo tiene una ruta 'api/cuit/{cuit}' donde trata de parsear el numero de cuit de una request del path {cuit} y así devolver de manera aleatoria un estado para ese cuit simulando la [api de central de deudores del banco centrar de argentina](http://www.bcra.gob.ar/BCRAyVos/Situacion_Crediticia.asp)

# Como ejecutar localmente

Primero debermos tener instalado go. En caso de no tenerlo podemos hacerlo siguien la [guia de instalación](https://golang.org/doc/install).

Ejecutar en la terminal el siguiente comando

```bash
go run main.go
```

Si todo salio bien veremos el mensaje

```
Ejecutando aplicación en el puerto 8080
```
