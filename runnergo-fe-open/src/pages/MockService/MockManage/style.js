import styled from '@emotion/styled';
import { css } from '@emotion/css';

export const ApisWrapper = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  .mock-response{
    .collapse-body{
      height: 500px !important;
    }
  }
  .apipost-collapse .apipost-collapse-item:not(:first-child) {
    border-top: none;
    }
  .apipost-collapse{
    width: 100%;
    height: 100%;
    background-color: var(--module);
    overflow: auto;
    .apipost-collapse-item{
      border-bottom: 2px solid var(--bg);
      .collapse-header{
        font-size: 14px;
        font-weight: 400;
        color: var(--content-color-primary);
        margin: 0 10px;
        padding: 0;
        svg{
          stroke:var(--content-color-primary);
        }
      }
    .collapse-body{
    height: 100%;
    border-top:2px solid var(--bg) !important;
    position: relative;
    margin: 0 10px;
    &>div:first-child{
      padding: 0 !important;
    }
    }
    }
  }
`;

export const ApisContentWarper = styled.div`
  flex: 1;
  height: 0;
  // background-color: var(--module);
`;

export const ApiHeaderWrapper = css`
  .api-manage-group {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 30px;
    padding: 1px;
    border-radius: var(--border-radius-default);
    border: 1px solid var(--border-color-strong);
    .enviroment-panel {
      border-left: 1px solid var(--border-color-strong);
      height: 100%;
    }
    .manage-item {
      display: flex;
      white-space: nowrap;
      padding: 0 6px;
      margin-left: 6px;
      height: 24px;
      line-height: 24px;
      border-radius: 3px;
      align-items: center;
      color: var(--content-color-secondary);
      cursor: pointer;
      svg {
        fill: var(--content-color-secondary);
      }
      &:hover {
        color: var(--content-color-primary);
        border-color: 0;
        background-color: var(--highlight-background-color-tertiary);
        svg {
          background-color: var(--highlight-background-color-tertiary);
          fill: var(--content-color-primary);
        }
      }
    }
    .bak-item,
    .bak-svg {
      display: flex;
      align-self: center;
      cursor: pointer;
      svg {
        background-color: var(--highlight-background-color-tertiary);
        fill: var(--content-color-primary);
      }
    }

    .apipost-btn {
      height: 24px;
    }
  }

  .btn-save {
    width: 84px;
    height: 24px;
    background: var(--log-blue);
    color: var(--font-1)fff;
    border-radius: 3px;
    display: flex;
    align-items: center;
    .left {
      height: 100%;
      padding: 0 8px;
      border-radius: 3px 0 0 3px;
      display: flex;
      align-items: center;
      justify-content: center;
      flex: 1;
      cursor: pointer;
    }
    .right {
      width: 24px;
      height: 24px;
      border-radius: 0 3px 3px 0;
      display: flex;
      justify-content: center;
      align-items: center;
      cursor: pointer;
      border-left: 1px solid rgba(0, 0, 0, 0.06);
    }
    svg {
      fill: var(--font-1)fff;
    }
  }
  @media screen and (max-width: 1400px) {
    .manage-item {
      & > span,
      .only-large {
        display: none;
      }
    }
  }
`;
