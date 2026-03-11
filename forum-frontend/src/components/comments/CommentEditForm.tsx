import React, { useState } from 'react';
import './Comment.css';

interface CommentEditFormProps {
  initialContent: string;
  onSave: (content: string) => void | Promise<void>;
  onCancel: () => void;
}

const CommentEditForm: React.FC<CommentEditFormProps> = ({ initialContent, onSave, onCancel }) => {
  const [content, setContent] = useState(initialContent);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!content.trim()) return;
    setLoading(true);
    await onSave(content);
    setLoading(false);
  };

  return (
    <form onSubmit={handleSubmit} style={{ marginTop: '12px' }}>
      <textarea
        value={content}
        onChange={e => setContent(e.target.value)}
        rows={3}
        style={{ width: '100%', marginBottom: '6px' }}
      />
      <div>
        <button type="submit" disabled={loading}>
          {loading ? 'Saving...' : 'Save'}
        </button>
        <button type="button" onClick={onCancel} style={{ marginLeft: '8px' }}>
          Cancel
        </button>
      </div>
    </form>
  );
};

export default CommentEditForm;