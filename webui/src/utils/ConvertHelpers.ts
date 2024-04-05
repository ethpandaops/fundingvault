
export function base64ToHex(str: string) {
  const raw = atob(str);
  let result = '';
  for (let i = 0; i < raw.length; i++) {
    const hex = raw.charCodeAt(i).toString(16);
    result += (hex.length === 2 ? hex : '0' + hex);
  }
  return result;
}

export function toDecimalUnit(amount: number | bigint, decimals?: number): number {
  let factor = Math.pow(10, decimals || 0);
  if(typeof amount === "bigint")
    amount = Number(amount);
  return amount / factor;
}

export function toBigintUnit(amount: number | bigint, decimals?: number): bigint {
  let factor = Math.pow(10, decimals || 0);
  if(typeof amount === "bigint")
    amount = Number(amount);
  return BigInt(amount * factor);
}

export function toReadableAmount(amount: number | bigint, decimals?: number, unit?: string, precision?: number): string {
  if(typeof decimals !== "number")
    decimals = 18;
  if(!precision) 
    precision = 3;
  if(!amount)
    return "0"+ (unit ? " " + unit : "");
  if(typeof amount === "bigint")
    amount = Number(amount);

  let decimalAmount = toDecimalUnit(amount, decimals);
  let precisionFactor = Math.pow(10, precision);
  let amountStr = (Math.round(decimalAmount * precisionFactor) / precisionFactor).toString();

  return amountStr + (unit ? " " + unit : "");
}

export function toReadableDuration(duration: number | bigint, maxParts?: number): string {
  if(typeof duration === "bigint")
    duration = Number(duration);
  if(typeof maxParts != "number")
    maxParts = 0;
  let res = "";
  let factor;
  let parts = 0;

  factor = (60 * 60 * 24 * 30);
  if(duration >= factor && (maxParts == 0 || parts < maxParts)) {
    let val = Math.floor(duration / factor);
    duration -= val * factor;
    res += (res ? " " : "") + val + "M";
    parts++;
  }

  factor = (60 * 60 * 24);
  if(duration >= factor && (maxParts == 0 || parts < maxParts)) {
    let val = Math.floor(duration / factor);
    duration -= val * factor;
    res += (res ? " " : "") + val + "d";
    parts++;
  }

  factor = (60 * 60);
  if(duration >= factor && (maxParts == 0 || parts < maxParts)) {
    let val = Math.floor(duration / factor);
    duration -= val * factor;
    res += (res ? " " : "") + val + "h";
    parts++;
  }

  factor = (60);
  if(duration >= factor && (maxParts == 0 || parts < maxParts)) {
    let val = Math.floor(duration / factor);
    duration -= val * factor;
    res += (res ? " " : "") + val + "min";
    parts++;
  }

  factor = 1;
  if(duration >= factor && (maxParts == 0 || parts < maxParts)) {
    let val = Math.floor(duration / factor);
    duration -= val * factor;
    res += (res ? " " : "") + val + "min";
    parts++;
  }

  return res;
}

