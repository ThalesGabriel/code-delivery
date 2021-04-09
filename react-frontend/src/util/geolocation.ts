export function getCurrentPosition(
    options?: PositionOptions
  ): Promise<{ lat: number; lng: number }> {
    return new Promise((resolve, reject) => {
      navigator.geolocation.getCurrentPosition(
        (positions) =>
          resolve({
            lat: positions.coords.latitude,
            lng: positions.coords.longitude,
          }),
        (error) => reject(error),
        options
      );
    });
  }