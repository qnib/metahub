FROM alpine

ARG INST
ARG SIZE
ARG HT
ENV INST=${INST}
ENV SIZE=${SIZE}
ENV HT=${HT}

COPY bin/entry.sh /usr/local/bin/
CMD ["/usr/local/bin/entry.sh"]
