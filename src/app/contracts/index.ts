// Sử dụng `any` một cách có kiểm soát để giải quyết vấn đề import JSON
import facadeAbiData from './DashboardFacade-abi.json';
import serviceStatusAbiData from './ServiceStatusContract-abi.json';
import recruitingAbiData from './RecruitingContract-abi.json';
import supportFundAbiData from './SupportFundContract-abi.json';
import timeSeriesAbiData from './TimeSeriesContract-abi.json';
import keyValueAbiData from './KeyValueContract-abi.json';


// --- Facade Contract (Dùng cho các lệnh gọi READ) ---
export const dashboardFacadeAddress = process.env.NEXT_PUBLIC_DASHBOARD_FACADE_ADDRESS as `0x${string}`;
export const dashboardFacadeAbi = facadeAbiData;


// --- Service Status Contract (Dùng để lắng nghe EVENT) ---
export const serviceStatusContractAddress = process.env.NEXT_PUBLIC_SERVICE_STATUS_CONTRACT_ADDRESS as `0x${string}`;
export const serviceStatusContractAbi = serviceStatusAbiData;


// --- Recruiting Contract (Dùng để lắng nghe EVENT) ---
export const recruitingContractAddress = process.env.NEXT_PUBLIC_RECRUITING_CONTRACT_ADDRESS as `0x${string}`;
export const recruitingContractAbi = recruitingAbiData;


// --- Support Fund Contract (Dùng để lắng nghe EVENT) ---
export const supportFundContractAddress = process.env.SUPPORT_FUND_CONTRACT_ADDRESS as `0x${string}`;
export const supportFundContractAbi = supportFundAbiData;


// --- Time Series Contract (Dùng để lắng nghe EVENT) ---
export const timeSeriesContractAddress = process.env.TIME_SERIES_CONTRACT_ADDRESS as `0x${string}`;
export const timeSeriesContractAbi = timeSeriesAbiData;


// --- Key Value Contract (Dùng để lắng nghe EVENT) ---
export const keyValueContractAddress = process.env.KEY_VALUE_CONTRACT_ADDRESS as `0x${string}`;
export const keyValueContractAbi = keyValueAbiData;