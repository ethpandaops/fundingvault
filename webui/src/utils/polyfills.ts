import { Buffer } from 'buffer';

window.global = window.global ?? window;
window.Buffer = window.Buffer ?? Buffer;
window.process = (window.process ?? { env: {} }) as any; // Minimal process polyfill

export {};
