  // Register Service Worker and sign up for push notifications
  (() => {
    if ('serviceWorker' in navigator) {
      navigator.serviceWorker.register('./assets/service-worker.js')
      .then(serviceWorkerRegistration => {
        console.log('[Main] Service Worker registered');
        serviceWorkerRegistration.pushManager.subscribe({
          userVisibleOnly: true
        }).then(function(subscription) {
          console.log('[Main] Subscribed to push notifications, endpoint:',
              subscription.endpoint.replace('https://android.googleapis.com/gcm/send/', ''));
        });
      }).catch(err => {
        console.log('[Main]', err);
      });
    }
  })();

