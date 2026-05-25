import React from 'react';

type TagItemProps = {
  tag: string;
  onRemove: (tag: string) => void;
};

function TagItem({ tag, onRemove }: TagItemProps) {
  return (
    <span className="tag-pill">
      <span>{tag}</span>
      <button
        type="button"
        className="tag-remove-button"
        onClick={() => onRemove(tag)}
        aria-label={`Remove ${tag}`}
      >
        ×
      </button>
    </span>
  );
}

export default TagItem;
