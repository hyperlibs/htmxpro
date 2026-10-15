// ../htmxpro/worker.ts
var worker_default = {
  async fetch(request, env) {
    return env.ASSETS.fetch(request);
  }
};
export {
  worker_default as default
};
