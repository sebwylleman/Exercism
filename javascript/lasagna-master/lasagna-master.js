/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export function cookingStatus(timer) {
  if (timer === 0) {
    return 'Lasagna is done.';
  }
  if (!timer) {
    return 'You forgot to set the timer.';
  }
  return 'Not done, please wait.';
}

export function preparationTime(layers, time = 2) {
  return layers.length * time;
}

export function quantities(layers) {
  // finish this
}
