FROM jrei/systemd-centos:7

COPY InnovationMinning.viper.yaml /etc/InnovationMinning/InnovationMinning.viper.yaml
COPY innovation-minning /usr/local/bin/innovation-minning

CMD ["innovation-minning", "run"]
