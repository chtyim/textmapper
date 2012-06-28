/**
 * Copyright (c) 2010-2012 Evgeny Gryaznov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textway.lapg.idea.editor;

import com.intellij.lang.Commenter;

/**
 * Gryaznov Evgeny, 1/31/11
 */
public class LapgCommenter implements Commenter {
	public String getLineCommentPrefix() {
		return "#";
	}

	public String getBlockCommentPrefix() {
		return null;
	}

	public String getBlockCommentSuffix() {
		return null;
	}

	public String getCommentedBlockCommentPrefix() {
		return null;
	}

	public String getCommentedBlockCommentSuffix() {
		return null;
	}
}