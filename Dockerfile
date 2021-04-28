FROM alpine
ADD wcb_micro_v2-service /wcb_micro_v2-service
ENTRYPOINT [ "/wcb_micro_v2-service" ]
