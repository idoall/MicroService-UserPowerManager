FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.role /mshk.role
RUN mkdir -p /logs

ENTRYPOINT [ "/mshk.role" ]
