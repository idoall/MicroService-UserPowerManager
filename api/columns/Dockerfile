FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.columns /mshk.columns
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.columns" ]
