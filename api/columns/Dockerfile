FROM alpine
ADD conf /conf
ADD mshk.columns /mshk.columns
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.columns" ]
