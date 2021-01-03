import { Octokit } from "@octokit/rest";
import config from "../data/config.json";

const { apiBaseUrl: baseUrl } = config["status-website"] || {};
const userAgent = config.userAgent;

export const createOctokit = () => {
  let token = "";
  if (
    typeof window !== "undefined" &&
    "localStorage" in window &&
    localStorage.getItem("personal-access-token")
  )
    token = localStorage.getItem("personal-access-token");
  return new Octokit({
    baseUrl,
    userAgent,
    auth: token || undefined,
  });
};

export const handleError = (error) => {
  if (error.message === "Bad credentials") {
    window.location.href = config.path + "/error";
  } else if ((error.message || "").indexOf("rate limit exceeded") > -1) {
    window.location.href = config.path + "/rate-limit-exceeded";
  } else {
    window.location.href = config.path + "/error";
    console.log(error.message);
  }
};

/**
 * Memoize a GitHub API response in local storage
 * @param {string} key - Local storage cache key
 * @param {Function} fn - Function that returns the result
 */
export const cachedResponse = async (key, fn) => {
  try {
    if (typeof window !== "undefined") {
      if ("localStorage" in window) {
        const data = localStorage.getItem(key);
        if (data) {
          const item = JSON.parse(data);
          if (
            new Date().getTime() - new Date(item.createdAt || "").getTime() >
            (document.domain === "localhost"
              ? config["status-website"].localhostCacheTime || 3600000
              : config["status-website"].productionCacheTime || 120000)
          ) {
            localStorage.removeItem(key);
          } else {
            console.log("Got cached item");
            return item.data;
          }
        }
      }
    }
  } catch (error) {}
  console.log("Got here");
  const i = await fn();
  localStorage.setItem(key, JSON.stringify({ data: i, createdAt: new Date() }));
  return i;
};
