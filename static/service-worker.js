// Self-destructing service worker.
//
// The previous version of this site registered a cache-first service worker
// at this exact URL, which would serve returning visitors the old site
// forever. Browsers re-fetch a registered service worker script on
// navigation (bypassing its own fetch handler), so shipping this file at the
// same URL heals those visitors: it wipes every cache, unregisters itself,
// and reloads open tabs. The new site registers its own worker at a
// different URL (app-worker.js). This file must stay deployed permanently.
self.addEventListener('install', () => self.skipWaiting());

self.addEventListener('activate', event => {
  event.waitUntil((async () => {
    const keys = await caches.keys();
    await Promise.all(keys.map(key => caches.delete(key)));
    await self.registration.unregister();
    const clients = await self.clients.matchAll({ type: 'window' });
    clients.forEach(client => client.navigate(client.url));
  })());
});
