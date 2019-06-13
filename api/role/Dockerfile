FROM alpine
ADD conf /conf
ADD mshk.role /mshk.role
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.role" ]
