FROM ghcr.io/renovatebot/renovate:38.110.1

ADD ./taskq_renovate /usr/bin/taskq_renovate

ENTRYPOINT taskq_renovate
