FROM alpine
ADD static /static
ADD conf /conf
ADD views /views
ADD mshk.web /mshk.web
WORKDIR /
ENTRYPOINT [ "/mshk.web" ]
