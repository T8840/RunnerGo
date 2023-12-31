import { css } from '@emotion/css';

export const DropWrapper = css`
  background-color: var(--background-color-primary);
  max-width: 267px;
  max-height: 260px;
  overflow: hidden;
  overflow-y: auto;
  .drop-item {
    display: flex;
    align-items: center;
    padding: 8px;
    border-radius: 0;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    cursor: pointer;
    &:hover {
      background-color: var(--select-hover);
    }
    svg {
      fill: var(--font-1);
      margin-right: 3px;
    }
  }
  &.backup {
    width: 100px;
  }
`;
